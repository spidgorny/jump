# jump &lt;folder name>

One of the most frustrating things when working in a terminal (CMD, bash) is to chdir to a directory.
Especially if this directory is not nearby like "..", or "../src".
One has to remember the current working directory (cwd) in memory or check it with "pwd" from time to time.

`jump` solves this problem by traversing the folder structure for you. 
You only need to specify the destination, not the path (!) to the destination.

## Example

Suppose we have the following directory structure:

```
/home/user/dev/
 - project1
     - bin
     - dest
     - src
        - components
        - views
        - config
        - hacks
     - node_modules
 - project2
    - bin
    - dest
    - src
       - components
       - views
       - config
       - hacks
    - node_modules
```

And the CWD is `/home/user/project1/src`. 

Then we can switch to *views* by `jump views` instead of `cd views`. This behavior is identical to *cd*.

Switchinig to *dest* folder can be done with `jump dest` instead of `cd ../dest`.

Going to *project2* is easy with `jump project2` instead of `cd ../../project2`.

In other words, `jump` will switch to the nearest folder matching the name you specified
 - inside the current folder or if nothing is found then
 - outside the current folder but nearest to *cwd*.

> We can also combine the folder and subfolder like this:
`jump project2/src`. (future idea not supported yet)

## Flags

 - -v for verbose mode

## Performance

Traversing the folder structure is not as fast as mere `cd` command.
But typing the valid argument for `cd` (even with autocomplete) also takes some time.
`jump` saves the typing time and mental effort for a price of traversing duration.

## Linux issues

`bash` does not allow the commands run from bash to make changes to the CWD.
Or it allows, but restores the CWD of the user after the program is finished.
To allow it to actually change the directory you need to import a small bash function which takes care of it for you.

`source jump`

Code is simple:

```
#!/bin/bash
function jump() {
  # output=$(go run src/jump.go $1)
  output=$(bin/jump $1)
  cd $output
  ls -l
}
```

Change it accordingly to the folder where `jump` is installed.

## Tech stack

Written in `go` to be fast. Tested on Windows, Linux and MacOS.
Bug reports and PR are welcome.

## Future plans

Jumping to a folder may take 5 seconds or more depending on the the distance to the destination from CWD.
Maybe we should cache the folder structure collected by the first execution of the `jump` 
and save it in JSON file somewhere to be reused on the next `jump` commands?

Adding this to a Linux repository will allow for an easy installation by everybody without installing golang compiler.
