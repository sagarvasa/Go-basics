# interface - summary

1. Interfaces helps in maintaining scalability in GO. They are "type" in go just like struct.

2. Stuct are basically data container since they hold data but interface is behviour container since it describe behaviour

3. We dont explicity use implements keyword for interface. its implicit to GO

4. We can create our own implementation of inbuild method also. (like Write method of I/O library). Any type can have method associated with it can be implemented using interface

5. Interface composition (interface of interfce)
    1. Smaller interface with less number of methods are suitable for building scalable applications
    2. for interface composition, we just have to mention interface name in another interface without datatype


6. Empty interface (marker interface)
special kind of interface which we can use with reflect libary to achieve more functionality
