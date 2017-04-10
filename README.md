# Snap Hawkular publisher plugin 
This plugin publishes snap metric data into Hawkular.

It's used in the [snap framework](http://github.com/intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Operating Systems](#operating-systems)
  * [Installation](#installation)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements
* [snap] (https://github.com/intelsdi-x/snap)
* [cassandra (3.7)](http://cassandra.apache.org)
* [golang 1.6+](https://golang.org/dl/)

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download cassandra publisher plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [Github Releases](https://github.com/intelsdi-x/snap/releases) page.

#### Building from source
* Get the package: 
```go get github.com/intelsdi-x/snap-plugin-publisher-hawkular```
* Build the snap-plugin-publisher-kawkular plugin
1. From the root of the snap-plugin-publisher-hawkular path type ```make all```.
* This builds the plugin in `/build/[GOOS]/[GOARCH]`.

#### Install Cassandra
* install Cassandra
```
 docker run -d --name snapcass -e CASSANDRA_START_RPC=true cassandra:3.0.9
```
* install Hawkular services
```
 docker run -d --name snaphawk --link=snapcass -e CASSANDRA_NODES=snapcass -p 8080:8080 -e ADMIN_TOKEN=topsecret hawkular/hawkular-services:latest
```

## Documentation

The plugin expects you to provide the following parameters:
 - `server` the hawkular server name or ip address.

You can also set the following options if needed:
 - `tenant` defaults to `snap`. It's required by hawkular.
 - `user` defaults to `jdoe` (string). 
 - `password` defaults to `password` (string).
 - `port` the hawkular server port. it defaults to `8080`.
 - `insecureSkipVerify` defaults to `true` (bool).
 - `scheme` defaults to `http` (string).

### Examples
See [examples/tasks](https://github.com/intelsdi-x/snap-plugin-publisher-hawkular/tree/master/examples/tasks) folder for examples.  

### Roadmap
This plugin is still in active development. As we launch this plugin, we have a few items in mind for the next few releases: 
 * Additional error handling
 * Large test
 
If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-publisher-hawkular/issues/new) and/or 
submit a [pull request](https://github.com/intelsdi-x/snap-plugin-publisher-hawkular/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap.

To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@candysmurf](https://github.com/candysmurf)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.

