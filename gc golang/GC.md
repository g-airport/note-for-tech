## Golang GC 

- v1.1 STW
- v1.3 Mark STW && Sweep Parallels
- v1.5 Tri-colour marking （Mark Sweep && STW）
- v1.8 hybrid write barrier


### Mark && Sweep

- phase 1 : mark
- phase 2 : sweep

        When GC Start, execute STW(Stop the world)
             0. Stop the world 
             1. Find process , check all reachable objects and mark them 
             2. Sweep unmark objects 
             3. Start the world

        Encounter Problem 
            1.Scan all heap for marking
            2.Sweep will generate heap fragments


        So Tri-colour marking Comming...
        
###  Tri-colour marking 

- phase 1 : The program create all objects marked white
- phase 2 : Scan all reachable objects marked gray (When GC Start)
- phase 3 : Mark objects of gray ( \
&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;step 1. like quote object marked gray \
&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;step 2. mark self with black \
&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;step 3. monitor memory updating && goto step 1 \
&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;step 4. when gray objects disappear, sweep white objects \
&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;step 5. finally , make black to white (current gc finished,next begin) \
) 

        Solve Problem 
            1.Sweep && User mode logic will make currency.
        Encounter Problem 
            1.User logic will update quote object.
            2.How to solve Process generating new object ?
            3.How to make mark and User mode logic currency ?
           


        So Write Barrier Comming...
        

### Write Barrier && Hybrid GC

    1.Write operator before && after compared.
    2.It is first perceived by other components of the system.
    (means gc processing:
        monitor memory update && remark object(stop the world)
    )
    3.When process generate new object,mark gray
    
 
    
    When Black Object quoted White Object
            1.Write Barrier trigger
            2.Send Signal to GC
            3.GC rescan objects && mark gray
            
    So, when GC start, whatever objects with created or quoted
    once updating, mark gray.