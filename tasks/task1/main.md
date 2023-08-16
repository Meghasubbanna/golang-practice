What is tricolor garbage collector?

Golang’s garbage collector uses a “tricolor” algorithm. 
This means it divides the heap objects into three sets: black, white, and grey. 
Initially, all objects are white, and as the algorithm proceeds, objects are moved into the grey and then black set.
In such a way that eventually the orphaned (collectible) objects are left in the white set, which is then cleared. 