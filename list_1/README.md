# Leader election algotihm simulator

Simulator of leder election algorithm which is presented (with analysis) [here](https://ki.pwr.edu.pl/lemiesz/AA/lider.pdf).

## How to run program

If you want to run this simulator, you need to install `Go` language compiler (with proper tools) in version > `1.11`, because of Go modules support.

To run program simply type in console `go run main.go`. You can use additional parameters described below:

* `-i` - number of iterations (repetitions) of algorithm _(default: 1000)_
* `-n` - number of nodes _(default: 1000)_
* `-u` - upper limit of nodes _(default: 1000)_
* `-t` - run prepared tests