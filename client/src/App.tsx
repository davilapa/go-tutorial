import "./App.css";
import { Container, Stack } from "@chakra-ui/react";
import Navbar from "./components/navbar";
import TodoForm from "./components/todo-form";
import TodoList from "./components/todo-list";

export const BASE_URL = import.meta.env.MODE === 'development' ? "http://localhost:4000/api/todos" : '/api';

function App() {
  return (
    <Stack h='100vh'>
      <Navbar />
      <Container maxW={"900px"}>
        <TodoForm />
        <TodoList />
      </Container>
    </Stack>
  );
}

export default App;
