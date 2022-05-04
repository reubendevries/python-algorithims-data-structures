# Data Structures

## Types of Data Structures
````mermaid
flowchart TB
    subgraph Data Structures
        subgraph Primitive
            A1-->string
            A2-->Interger
            A3-->Boolean
            A4-->Float
            A5-->Character
        end
        subgraph Non-Primitive
            subgraph Linear
                subgraph Static
                    B1-->Array
                end
                subgraph Dynamic
                    C1-->Linked List
                    C2-->Stack
                    C3-->Queue
                end
            end
            subgraph Non-Linear
                D1-->Tree
                D2-->Graph
            end
    end
````