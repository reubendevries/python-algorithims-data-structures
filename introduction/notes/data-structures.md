# Data Structures

## Types of Data Structures
````mermaid
flowchart TB
    subgraph Data Structures
        direction LR
        subgraph Primitive
            (String)
            (Interger)
            (Boolean)
            (Float)
            (Character)
        end
        subgraph Non-Primitive
            subgraph Linear
                direction LR
                subgraph Static
                    (Array)
                end
                subgraph Dynamic
                    (Linked List)
                    (Stack)
                    (Queue)
                end
            end
            subgraph Non-Linear
                (Tree)
                (Graph)
            end
    end
````