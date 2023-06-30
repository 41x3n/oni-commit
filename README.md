# oni-commit

Commit Carefully

Ideally I suggest using [semantic-git-commit-cli](https://www.npmjs.com/package/semantic-git-commit-cli) but I wanted to learn how that package works while learning Golang side by side.

![Demo Gif](https://imgur.com/Va8ynPP.gif)

### Requirements -

1. Git
2. Vim

### Installation -

1. MacOs -
    1. `go build -o build/oni`
    2. `sudo mv ./build/oni /usr/local/bin/`
    3. Restart terminal.
    4. In any git repository folder, type `oni` and press enter.

    <br>

2. Windows -
    1. `go build -o build/oni.exe` 
    2. Move the executable file to any folder. Example : `C:\<Any Folder>\oni.exe`
    3. Copy path to that folder. Example : `C:\<Any Folder>\oni`
    4. Add the path to the environment variable.
    5. Restart terminal.
    6. In any git repository folder, type `oni` and press enter.