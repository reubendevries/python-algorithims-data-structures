# Data Structures

## Types of Data Structures
````mermaid
flowchart TB
    subgraph Data_Structures
        direction TB
        subgraph Primitive
            String
            Interger
            Boolean
            Character
            Float
        end
        subgraph Non-Primitive
            direction TB
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
                Tree
                Graph
            end
        end
    end
    Data_Structures --> Primitive
    Data_Structures --> Non-Primitive
    Non-Primitive --> Linear
    Non-Primitive --> Non-Linear
    Linear --> Static
    Linear --> Dynamic
````