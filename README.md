# goTraining

Results of taking several courses from Jon Calhoun via calhoun.io. First foray into a new language, and really coding, since 2013. Built several projects using a test-first approach with the help of the excellent test facilities provided in Go.

Modules include: 

(1) a quiz app that takes in data from a file and runs a timed "game" where a user answers questions from a file. Uses a go func to run the timer in the background and interrupt the main thread to end the game

(2) a urlshortener a-la bit.ly that redirects based on a data file which can be key/value pairs (Map), yaml, or json formats. Stored as either a server file or in BoltDB

(3) choose your own adventure (cyoa) app that builds a story as a tree from a json file and either displays it to the console or a simple web page. 

(4) link parser which reads html elements and stores all of the links from the a href tags, ignoring nested links

(5) sitemap builder using the project from (4) that parses links within a domain and exports the resultant sitemap to xml

(6) implementations of caesar cipher and camel count challenges

(7) task manager with persistent storage implemented with a command line interface using the spf13/cobra clt package

(8) a phone number normalizer that takes in various phone numbers (ostensibly from signups at a conference) and normalizes them using regular expressions

(9) a clone of HackerNews that used concurrency to pull stories in parallel instead of serially (while maintaining the original order)

(10) an tool that produces simple svg images including bar graphs and crazy looking fractals 

Really enjoyed what I learned and took me back to my C++/Java days in college. Many thanks to Jon for the course!
