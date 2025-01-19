import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";

import TaskItem from "./TaskItem";
import { useQuery } from "@tanstack/react-query";
import { BASE_URL } from "../App";

export type Task = {
	_id: number;
	body: string;
	completed: boolean;
};

const TaskList = () => {
	const { data: tasks, isLoading } = useQuery<Task[]>({
		queryKey: ["tasks"],
		queryFn: async () => {
			try {
				const res = await fetch(BASE_URL + "/tasks");
				const data = await res.json();

				if (!res.ok) {
					throw new Error(data.error || "Something went wrong");
				}
				return data || [];
			} catch (error) {
				console.log(error);
			}
		},
	});

	return (
		<>
			<Text
				fontSize={"4xl"}
				textTransform={"uppercase"}
				fontWeight={"bold"}
				textAlign={"center"}
				my={2}
				bgGradient='linear(to-l, #0b85f8, #00ffff)'
				bgClip='text'
			>
				Today's Tasks
			</Text>
			{isLoading && (
				<Flex justifyContent={"center"} my={4}>
					<Spinner size={"xl"} />
				</Flex>
			)}
			{!isLoading && tasks?.length === 0 && (
				<Stack alignItems={"center"} gap='3'>
					<Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
						All tasks completed! 🤞
					</Text>
					<img src='/go.png' alt='Go logo' width={70} height={70} />
				</Stack>
			)}
			<Stack gap={3}>
				{tasks?.map((task) => (
					<TaskItem key={task._id} task={task} />
				))}
			</Stack>
		</>
	);
};
export default TaskList;

// STARTER CODE:

// import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
// import { useState } from "react";
// import TaskItem from "./TaskItem";

// const TaskList = () => {
// 	const [isLoading, setIsLoading] = useState(true);
// 	const Tasks = [
// 		{
// 			_id: 1,
// 			body: "Buy groceries",
// 			completed: true,
// 		},
// 		{
// 			_id: 2,
// 			body: "Walk the dog",
// 			completed: false,
// 		},
// 		{
// 			_id: 3,
// 			body: "Do laundry",
// 			completed: false,
// 		},
// 		{
// 			_id: 4,
// 			body: "Cook dinner",
// 			completed: true,
// 		},
// 	];
// 	return (
// 		<>
// 			<Text fontSize={"4xl"} textTransform={"uppercase"} fontWeight={"bold"} textAlign={"center"} my={2}>
// 				Today's Tasks
// 			</Text>
// 			{isLoading && (
// 				<Flex justifyContent={"center"} my={4}>
// 					<Spinner size={"xl"} />
// 				</Flex>
// 			)}
// 			{!isLoading && Tasks?.length === 0 && (
// 				<Stack alignItems={"center"} gap='3'>
// 					<Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
// 						All tasks completed! 🤞
// 					</Text>
// 					<img src='/go.png' alt='Go logo' width={70} height={70} />
// 				</Stack>
// 			)}
// 			<Stack gap={3}>
// 				{Tasks?.map((Task) => (
// 					<TaskItem key={Task._id} Task={Task} />
// 				))}
// 			</Stack>
// 		</>
// 	);
// };
// export default TaskList;