![NSQustodian Logo](docs/nsqustodian-logo-transparent.png)

NSQustodian is a CLI tool for administrating and managing NSQ clusters. It allows you to context switch between multiple clusters and perform various administrative tasks, such as listing topics and channels, pausing and unpausing topics and channels, emptying message queues, listing NSQ nodes, and offloading and loading topics and channels to and from S3-compatible object storage.

Installation
------------

To build the NSQustodian binary from source, follow these steps:

1. Clone the NSQustodian repository:

```
git clone https://github.com/acjohnson/nsqustodian.git
```

2. Change to the NSQustodian directory:

```
cd nsqustodian
```

3. Build the NSQustodian binary:

```
go build
```

This will create an `nsqustodian` binary in the current directory.

To install the NSQustodian binary in the `~/bin` directory and add that directory to your PATH, follow these steps:

1. Copy the `nsqustodian` binary to the `~/bin` directory (you may need to create this directory if it does not already exist):

```
cp nsqustodian ~/bin
```

2. Add the `~/bin` directory to your PATH. You can do this by adding the following line to your shell configuration file (e.g. `~/.bashrc` or `~/.zshrc`):

```
export PATH=$PATH:$HOME/bin
```

3. Reload your shell configuration file to apply the changes:

```
source ~/.bashrc
```

or

```
source ~/.zshrc
```

After following these steps, you should be able to run the `nsqustodian` command from anywhere on your system.

For more information on working with the Go toolchain and building Go programs, please see the [Go documentation](https://golang.org/doc/).

Configuration
-------------

Before you can use NSQustodian, you need to create a configuration file. You can do this by running the following command:

```
nsqustodian config create-context --name my-app-nsq --nsq-admin nsqadmin.example.com
```

This will create a configuration file at `~/.nsqustodian.yaml` with the specified name and NSQ admin address.

Commands
--------

Once you have created a configuration file, you can use the following commands to administer your NSQ cluster:

* `nsqustodian channels list`: Lists all channels in the current context.
* `nsqustodian channels pause-channel --topic my-app-topic --channel my-app-channel`: Pauses the specified channel.
* `nsqustodian channels unpause-channel --topic my-app-topic --channel my-app-channel`: Unpauses the specified channel.
* `nsqustodian channels offload-channel --topic my-app-topic --channel my-app-channel --s3-bucket-name my-bucket --s3-bucket-key my-folder`: Offloads the messages in the specified channel to the specified S3 bucket.
* `nsqustodian topics list`: Lists all topics in the current context.
* `nsqustodian topics pause-topic --topic my-app-topic`: Pauses the specified topic.
* `nsqustodian topics unpause-topic --topic my-app-topic`: Unpauses the specified topic.
* `nsqustodian topics offload-topic --topic my-app-topic --s3-bucket-name my-bucket --s3-bucket-key my-folder`: Offloads the messages in the specified topic to the specified S3 bucket.

You can switch between different contexts by using the `config use-context` sub-command. For example:

```
nsqustodian config use-context --name my-other-nsq-context
```

Advanced Features
-----------------

NSQustodian also includes some advanced features for working with NSQ topics and channels. These features include:

* Loading and offloading topics and channels from S3-compatible object storage.

For more information on how to use these features, please see the `nsqustodian help` command.

License
-------

NSQustodian is licensed under the GPLv2. For more information, please see the [LICENSE](LICENSE) file.

Support
-------

If you have any questions or issues with NSQustodian, please open an issue on the [GitHub repository](https://github.com/acjohnson/nsqustodian).

Acknowledgments
--------------

NSQustodian was inspired by the [NSQ Admin](https://github.com/nsqio/nsq/tree/master/nsqadmin) web interface.

Thanks to the maintainers of those projects for their excellent work.
