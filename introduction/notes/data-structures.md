# Data Structures

## Types of Data Structures
````mermaid
flowchart LR
    subgraph Data Structures
        direction TB
        subgraph Primitive
            String
            Interger
            Boolean
            Character
            Float
        end
        subgraph Non-Primitive
            subgraph Linear
                direction TB
                subgraph Static
                    Array
                end
                subgraph Dynamic
                    Linked_List
                    Stack
                    Queue
                end
            end
            subgraph Non-Linear
                direction TB
                Tree
                Graph
            end
        end
    end
````