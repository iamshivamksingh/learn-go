1. Use channels when passing ownership of data
2. Use mutexes for managing state


Don't use embedding because it's convenient
1. Think about the effect embedding has on your public API.
2. Do you really want to expose these methods and have people coupling their own code to them?
3. With respect to mutexes, this could be potentially disastrous in very unpredictable and weird ways, imagine some nefarious code unlocking a mutex when it shouldn't be; this would cause some very strange bugs that will be hard to track down.