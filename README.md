# To Do CLI

This is a simple todo CLI, this code is based in [Golang Tutorial: Build A Beautiful CLI Todo App With Support for Piping](https://www.youtube.com/watch?v=j1CXoOQXbco)

![todo](./docs/todo.png)

## Requisites

First of all, [download](https://go.dev/dl/) and install Go. 1.20 or higher is required.

## Building and Running

First you need to build the project.

```bash
make build
```

After build, you have the project binaries and you can run the todo app with the commands:

### Add a task

```bash
./todo -add <your_task_name>
```

### List all tasks

```bash
./todo -list
```

### Complete a task

```bash
./todo -complete=<index>
```

### Delete a task

``` bash
./todo -del=<index>
./todo -del=1
```

## License

Licensed under the MIT License. See [`LICENSE`](LICENSE) for more details.
