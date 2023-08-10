import { useEffect, useState } from "react";
import { getAllTasks } from "../api/tasks.api";

function TasksList() {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    async function loadTask() {
      const res = await getAllTasks();

      setTasks(res.data);
    }

    loadTask();
  }, []);

  return (
    <div>
      <h1>Tasks List</h1>
      {tasks.map((task) => (
        <div key={task.id}>
          <h1>{task.title}</h1>
          <p>{task.description}</p>
        </div>
      ))}
    </div>
  );
}

export default TasksList;
