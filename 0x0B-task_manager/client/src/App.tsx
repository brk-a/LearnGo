import { Container, Stack } from "@chakra-ui/react";
import Navbar from "./components/Navbar";
import TaskForm from "./components/TaskForm";
import TaskList from "./components/TaskList";

export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:8080/api/v1" : "/api";

function App() {
	return (
		<Stack h='100vh'>
			<Navbar />
			<Container>
				<TaskForm />
				<TaskList />
			</Container>
		</Stack>
	);
}

export default App;