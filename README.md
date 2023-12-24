# ZIPPER

zipper is a CLI application to **ZIP** and **UNZIP** file written in Go.

## USAGE

### ZIPPING

**Flags:**  
`-z` To zipping  
`-fl` For zipping a file or files  
`-fd` For zipping a folder

```
$ zipper -z -type <file/folder-path> <destination-path>
```

If you want specify a **multiple** files; wrap it with `[]` and separate it by _space_.  
If you don't specify a destination path; your current directory will be the destination path and will be named zipper.zip.

#### EXAMPLE

```
$ zipper -z -fl [task1.pdf task2.pdf] homework
```

### UNZIPPING

**Flags:**  
`-u` To unzipping

```
$ zipper -u <zip-file-path> <destination-path>
```

Again, if you don't specify a destination path; your current directory will be the destination path inside folder named zipper.

#### EXAMPLE

```
$ zipper -u homework.zip homework-done
```
