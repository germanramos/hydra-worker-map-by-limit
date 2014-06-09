hydra-worker-map-by-limit
===========================

Worker for Hydra v3.  
Given an attribute and a value will generate two maps with the given value as separator.

# Installation

## Ubuntu/Debian

Add PPAs for:  
https://launchpad.net/~chris-lea/+archive/libpgm  
https://launchpad.net/~chris-lea/+archive/zeromq  
  
and run:  
```
sudo dpkg -i hydra-worker-map-by-limit-1-0.x86_64.deb
sudo apt-get install -f
```
## CentOS/RedHat/Fedora
```
sudo yum install libzmq3-3.2.2-13.1.x86_64.rpm hydra-worker-map-by-limit-1-0.x86_64.rpm
```

# Configuration

In apps.json:

- Name: "MapByLimit"
- Arguments:
  - limitAttr: The name of the numerical attribute that we are going compare
  - limitValue: the value to separate the maps
  - mapSort: normal=lower before, reverse=higher before

## Configuration example
```
{
	"worker": "MapByLimit",
	"limitAttr": "cpuLoad",
	"limitValue": 90,
	"mapSort": "normal"
}
```			
This will generate two maps. The first map with all the instances with cpuLoad lower than 90 and the second map with the instances with cpuLoad higher than 90.

## Service configuration

No additional configuration is needed if running in the same machine that Hydra.  
Tune start file at /etc/init.d/hydra-worker-map-by-limit if you run it in a separate machine.

# Run
```
sudo /etc/init.d/hydra-worker-map-by-limit start
```

# License

(The MIT License)

Authors:  
Germán Ramos &lt;german.ramos@gmail.com&gt;  
Pascual de Juan &lt;pascual.dejuan@gmail.com&gt;  
Jonas da Cruz &lt;unlogic@gmail.com&gt;  
Luis Mesas &lt;luismesas@gmail.com&gt;  
Alejandro Penedo &lt;icedfiend@gmail.com&gt;  
Jose María San José &lt;josem.sanjose@gmail.com&gt;  

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
