# adventure

Engine for writing and interpreting text-based adventure games.

## Features

### Story

Create your game in a declarative yaml-based format.

A game consists of states and actions. A start node also has to be specified,
where the game will begin ('start\_node').

#### Actions

An action can do the following things (in any combination):

1) Redirect to another state

2) Show a text associated with the action

3) Prompt for a 'secret', which is a required input by the user to 'unlock' the
action.

#### Example

```yml
---
version: "0.0.1 (development)"
start_node: startnode

entries:
   - name: startnode
     story: "This is a text which should appear at the beginning of the game"
     actions:
         - target: second_state
           name: "Step to the second state"
         - target: third_state
           name: "Name of the action to the third state"
   - name: second_state
     story: "Story of second state"
     actions:
         - target: startnode
           name: "move to the start node"
         - name: "Do something and stay still"
           story: "This text will appear if the user executed this action, which won't redirect to a new state."
           secret: "titok1234"
   - name: third_point
     story: "Story of the third and final point"
     is_end: true
```



