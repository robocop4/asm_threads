# asm_threads
![index](https://raw.githubusercontent.com/notrobot1/asm_threads/main/img/index.gif)

<p>This repository publishes the source code of the program to track the sequence of program execution. This program will be useful for people who are involved in reverse software development. The program will allow you to track the function after which the program execution follows another thread. <a href="https://rada.re/">Radare2</a> must be installed on your computer to work correctly. </p>

<p>Download the repository:</p>

`git clone https://github.com/notrobot1/asm_threads.git`

<p>Go to the go_cli directory and build the program.</p> 

`cd go_cli`

`go build main.go createJson.go` 

<p>After building, the main executable file will appear in the directory. To run go_cli, run  `./main -p /bin/ls -arg -la`  in the terminal. </p>
<p>The result will be two files in the current directory: `base.json` and `edges1.json`. </p>
<p>Let's run ls without arguments and see the differences between function calls. `./main -p /bin/ls -st` </p> 
<p>The result will be a new file in the current directory is  `edges2.json`.</p> 
<p>Let's go to the UI at the <a href="https://notrobot1.github.io/asm_threads/ui/build/">link</a>.  </p>
<p>Let's upload two files via the file submission form base.json and edges2.json. After uploading you will see the nodes and threads that connect them. Using the `next` button you can take steps. The red thread shows the chain of ls command calls with arguments and the blue thread shows the execution without arguments. If the program was executed along the same route, the red and blue colors overlap.</p>


