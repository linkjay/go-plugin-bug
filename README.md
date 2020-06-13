# Go Plugin Bug

This project is designed to illustrate how you can make plugins that work properly that work within Go. What's really cool about this is you can make dynamic loaders that load a bunch of Go scripts within a dir.

The problem is that Go needs to be compiled. This will work 100% okay if you run `build.sh` because it assumes the same Go compiler. However, if you use an IDE like VSCODE, it can call `build.sh` which will build everything, but it needs to run it's own version of Go to compile the main executable so that it can run in debug mode.

Thus is where the problem is. You tell your IDE to run the bash script which uses your system's Go to compile the plugins. Now your IDE will use its own version of Go to compile the main executable and run it in debug mode. This is where you get an error like this:

`2020/06/12 20:56:03 Fatal error when opening plugin! Err: plugin.Open("plugins/test-plugin"): plugin was built with a different version of package runtime/internal/sys`

# Usage

Run `build.sh` and it will compile the plugins, and the main executable. You can then run `./main` which is the main exectuable. That should output the code from `test-plugin.go` which just prints `hello`. 

Try now in your IDE. Setup a new task to run `main.go` in debug mode. Set it up so that it runs `build.sh` which will compile your plugins first. It will also create a main binary which could possibly used by your IDE, but isn't by default. You should then get the error above.