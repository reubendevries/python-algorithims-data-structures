# Data Structures

## Types of Data Structures
````mermaid
flowchart TB
    subgraph Data Structures
        direction LR
        subgraph Primitive
            directiion TB
            (String)
            (Interger)
            (Boolean)
            (Float)
            (Character)
        end
        subgraph Non-Primitive
            direction TB
            subgraph Linear
                direction LR
                subgraph Static
                    direction TB
                    (Array)
                end
                subgraph Dynamic
                    direction TB
                    (Linked List)
                    (Stack)
                    (Queue)
                end
            end
            subgraph Non-Linear
                direction TB
                (Tree)
                (Graph)
            end
    end
````