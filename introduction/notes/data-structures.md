# Data Structures

## Types of Data Structures
````mermaid
flowchart LR
    subgraph Data Structures
        direction LR
        subgraph Primitive
            direction TB
            String
            Interger
            Boolean
            Character
            Float
        end
        subgraph Non-Primitive
            direction TB
            subgraph Linear
                direction LR
                subgraph Static
                    direction TB
                    Array
                end
                subgraph Dynamic
                    direction TB
                    Linked_List
                    Stack
                    Queue
                end
            end
            subgraph Non-Linear
                direction LB
                Tree
                Graph
            end
        end
    end
````