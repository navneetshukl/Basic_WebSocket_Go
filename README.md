# WebSocket Communication with Gorilla WebSocket in Go

This repository demonstrates WebSocket communication using the Gorilla WebSocket package in Go. It includes a simple server-side code snippet and an explanation of the message flow between the client and server.

## Code Snippet Explanation

The provided code snippet showcases the usage of `ReadMessage` and `WriteMessage` methods in the Gorilla WebSocket package.

### ReadMessage

The `ReadMessage` method is used to read a message from the WebSocket connection. It returns three values:

- `msgType`: An integer representing the type of the received message, such as text, binary, or close message.
- `msg`: A byte array containing the actual message data.
- `err`: An error object that indicates whether there was an error while reading the message.

In the code snippet, the received message is printed to the console using `fmt.Printf` with the format `%s sent: %s\n`. It displays the remote address of the WebSocket client (`conn.RemoteAddr()`) and the content of the received message (`string(msg)`).

### WriteMessage

The `WriteMessage` method is used to write a message to the WebSocket connection. It takes two arguments:

- `msgType`: An integer representing the type of the message being sent, such as text, binary, or close message.
- `msg`: A byte array containing the message data to be sent.

The method sends the specified message to the WebSocket client. In the provided code snippet, the received message is immediately sent back to the client using `WriteMessage` to create an echo effect. If there is an error while writing the message, the function returns, which effectively ends the loop and closes the connection.

## Message Flow

The message flow between the client and server takes place over the WebSocket connection. Here's an overview of each step in the process:

1. **Client establishes a WebSocket connection with the server:**
   - The client initiates a WebSocket handshake by sending an HTTP request to the server.
   - The server responds with an HTTP response indicating whether the WebSocket handshake is successful.
   - Once the handshake is successful, the WebSocket connection is established, and both the client and server can communicate over it.

2. **Client sends a message to the server using the WebSocket connection:**
   - On the client side, JavaScript or a WebSocket client library is used to send messages over the WebSocket connection.
   - The client can send a message by calling the `send` method on the WebSocket object and passing the desired message as an argument.

3. **Server receives the message using the `ReadMessage` method:**
   - On the server side, a WebSocket handler is set up to handle incoming WebSocket connections and messages.
   - When the client sends a message over the WebSocket connection, the server's WebSocket handler is triggered.
   - Inside the WebSocket handler, the server uses the `ReadMessage` method to read the received message from the WebSocket connection.
   - The `ReadMessage` method returns three values: `msgType`, `msg`, and `err`.
     - `msgType` is an integer representing the type of the received message (e.g., text, binary, or close message).
     - `msg` is a byte array containing the actual message data.
     - `err` is an error object that indicates whether there was an error while reading the message.

4. **Server processes the received message:**
   - After reading the message using the `ReadMessage` method, the server can perform any necessary processing or logic based on the received message.
   - In the provided code snippet, the server simply prints the received message to the console using `fmt.Printf`, creating an echo effect.

5. **Server sends a response message back to the client using the `WriteMessage` method:**
   - After processing the received message, the server can send a response message back to the client.
   - The server uses the `WriteMessage` method to write the response message to the WebSocket connection.
   - The `WriteMessage` method takes two arguments: `msgType` and `msg`.
     - `msgType` is an integer representing the type of the message being sent (e.g., text, binary, or close message).
     - `msg` is a byte array containing the message data to be sent.
   - In the provided code snippet, the server immediately sends the received message back to the client as a response, creating an echo effect.
   - If there is an error while writing the message, the function returns, effectively ending the loop and closing the connection.

6. **Client receives the response message and handles it:**
   - On the client side, if the WebSocket connection is properly set up with appropriate event handlers, the client can receive the response message from the server.
   - The client can handle the received message by implementing the `onmessage` event handler.
   - Inside the `onmessage` event handler, the client can access the received message using the `event.data` property and perform any desired actions based on the content of the message.

Please refer to the code snippet and explanation for more details.

