package db

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value TaskData
}
type TaskData struct {
	Value      string
	Complete   bool
	UpdateTime time.Time
}

func main() {

}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func AddTask(task string) (int, error) {
	taskID := -1
	err := db.Update(func(tx *bolt.Tx) error {
		b := getTaskBucket(tx)
		id, _ := b.NextSequence()
		taskID = int(id)
		idbs := itob(id)
		td := TaskData{task, false, time.Now()}
		jsondata, jerr := json.Marshal(td)
		checkJsonError(jerr)
		return b.Put(idbs, jsondata)
	})
	return taskID, err
}

func GetIncompleteTasks() ([]Task, error) {
	tasks := make([]Task, 0)
	err := db.View(func(tx *bolt.Tx) error {
		b := getTaskBucket(tx)
		errin := b.ForEach(func(k, v []byte) error {
			key, err := btoi(k)
			td := TaskData{}
			jerr := json.Unmarshal(v, &td)
			checkJsonError(jerr)
			if !td.Complete {
				tasks = append(tasks, Task{int(key), td})
			}
			return err
		})
		return errin
	})
	return tasks, err
}

func GetCompletedTasks() ([]Task, error) {
	tasks := make([]Task, 0)
	err := db.View(func(tx *bolt.Tx) error {
		b := getTaskBucket(tx)
		errin := b.ForEach(func(k, v []byte) error {
			key, err := btoi(k)
			td := TaskData{}
			jerr := json.Unmarshal(v, &td)
			checkJsonError(jerr)
			if td.Complete {
				tasks = append(tasks, Task{int(key), td})
			}
			return err
		})
		return errin
	})
	return tasks, err
}

func RemoveTasks(taskIDs []int) ([]int, error) {
	taskmap := make(map[int]bool, len(taskIDs))
	removed := make([]int, 0)
	for _, id := range taskIDs {
		taskmap[id] = true
	}
	err := db.Update(func(tx *bolt.Tx) error {
		b := getTaskBucket(tx)
		ix := 0
		return b.ForEach(func(k, v []byte) error {
			if taskmap[ix+1] {
				fmt.Println("deleting", (ix + 1), "with key", string(k))
				err := b.Delete(k)
				if err != nil {
					return err
				}
				removed = append(removed, ix+1)
			}
			ix++
			return nil
		})
	})
	return removed, err
}

func CompleteTasks(taskIDs []int) ([]int, error) {
	taskmap := make(map[int]bool, len(taskIDs))
	completed := make([]int, 0)
	for _, id := range taskIDs {
		taskmap[id] = true
	}
	err := db.Update(func(tx *bolt.Tx) error {
		b := getTaskBucket(tx)
		ix := 0
		return b.ForEach(func(k, v []byte) error {
			if taskmap[ix+1] {
				td := TaskData{}
				jerr := json.Unmarshal(v, &td)
				checkJsonError(jerr)
				td.Complete = true
				newbytes, jerr := json.Marshal(td)
				checkJsonError(jerr)
				b.Put(k, newbytes)
				completed = append(completed, ix+1)
			}
			ix++
			return nil
		})
	})
	return completed, err
}

func itob(v uint64) []byte {
	ret := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(ret, v)
	return ret
}
func btoi(b []byte) (uint64, error) {
	return binary.ReadUvarint(bytes.NewReader(b))
}

func getTaskBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket(taskBucket)
}

func checkJsonError(jerr error) {
	if jerr != nil {
		fmt.Println("json error", jerr)
	}
}
