import { Badge, Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { Task } from "./TaskList";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { BASE_URL } from "../App";

const TaskItem = ({ task }: { task: Task }) => {
	const queryClient = useQueryClient();

	const { mutate: updateTask, isPending: isUpdating } = useMutation({
		mutationKey: ["updateTask"],
		mutationFn: async () => {
			if (task.completed) return alert("Task is already completed");
			try {
				const res = await fetch(BASE_URL + `/tasks/${task._id}`, {
					method: "PATCH",
				});
				const data = await res.json();
				if (!res.ok) {
					throw new Error(data.error || "Something went wrong");
				}
				return data;
			} catch (error) {
				console.log(error);
			}
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["Tasks"] });
		},
	});

	const { mutate: deleteTask, isPending: isDeleting } = useMutation({
		mutationKey: ["deleteTask"],
		mutationFn: async () => {
			try {
				const res = await fetch(BASE_URL + `/Tasks/${task._id}`, {
					method: "DELETE",
				});
				const data = await res.json();
				if (!res.ok) {
					throw new Error(data.error || "Something went wrong");
				}
				return data;
			} catch (error) {
				console.log(error);
			}
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["Tasks"] });
		},
	});

	return (
		<Flex gap={2} alignItems={"center"}>
			<Flex
				flex={1}
				alignItems={"center"}
				border={"1px"}
				borderColor={"gray.600"}
				p={2}
				borderRadius={"lg"}
				justifyContent={"space-between"}
			>
				<Text
					color={task.completed ? "green.200" : "yellow.100"}
					textDecoration={task.completed ? "line-through" : "none"}
				>
					{task.body}
				</Text>
				{task.completed && (
					<Badge ml='1' colorScheme='green'>
						Done
					</Badge>
				)}
				{!task.completed && (
					<Badge ml='1' colorScheme='yellow'>
						In Progress
					</Badge>
				)}
			</Flex>
			<Flex gap={2} alignItems={"center"}>
				<Box color={"green.500"} cursor={"pointer"} onClick={() => updateTask()}>
					{!isUpdating && <FaCheckCircle size={20} />}
					{isUpdating && <Spinner size={"sm"} />}
				</Box>
				<Box color={"red.500"} cursor={"pointer"} onClick={() => deleteTask()}>
					{!isDeleting && <MdDelete size={25} />}
					{isDeleting && <Spinner size={"sm"} />}
				</Box>
			</Flex>
		</Flex>
	);
};
export default TaskItem;

// STARTER CODE:

// import { Badge, Box, Flex, Text } from "@chakra-ui/react";
// import { FaCheckCircle } from "react-icons/fa";
// import { MdDelete } from "react-icons/md";

// const TaskItem = ({ Task }: { Task: any }) => {
// 	return (
// 		<Flex gap={2} alignItems={"center"}>
// 			<Flex
// 				flex={1}
// 				alignItems={"center"}
// 				border={"1px"}
// 				borderColor={"gray.600"}
// 				p={2}
// 				borderRadius={"lg"}
// 				justifyContent={"space-between"}
// 			>
// 				<Text
// 					color={Task.completed ? "green.200" : "yellow.100"}
// 					textDecoration={Task.completed ? "line-through" : "none"}
// 				>
// 					{Task.body}
// 				</Text>
// 				{Task.completed && (
// 					<Badge ml='1' colorScheme='green'>
// 						Done
// 					</Badge>
// 				)}
// 				{!Task.completed && (
// 					<Badge ml='1' colorScheme='yellow'>
// 						In Progress
// 					</Badge>
// 				)}
// 			</Flex>
// 			<Flex gap={2} alignItems={"center"}>
// 				<Box color={"green.500"} cursor={"pointer"}>
// 					<FaCheckCircle size={20} />
// 				</Box>
// 				<Box color={"red.500"} cursor={"pointer"}>
// 					<MdDelete size={25} />
// 				</Box>
// 			</Flex>
// 		</Flex>
// 	);
// };
// export default TaskItem;