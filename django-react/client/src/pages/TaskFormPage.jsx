import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import { createTask, deleteTask, getTask, updateTask } from "../api/tasks.api";
import { toast } from "react-hot-toast";

function TaskFormPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
  } = useForm();

  const navigate = useNavigate();
  const params = useParams();

  const onSubmit = handleSubmit(async (data) => {
    if (params.id) {
      await updateTask(params.id, data);
      toast.success("Task updated", {
        position: "bottom-right",
        style: {
          background: "#101010",
          color: "#fff",
        },
      });
    } else {
      await createTask(data);
      toast.success("New Task Added", {
        position: "bottom-right",
        style: {
          background: "#101010",
          color: "#fff",
        },
      });
    }

    navigate("/tasks");
  });

  useEffect(() => {
    async function loadTask() {
      if (params.id) {
        const { data } = await getTask(params.id);
        setValue("title", data.title);
        setValue("description", data.description);
      }
    }
    loadTask();
  }, [params.id, setValue]);

  return (
    <div className="max-w-xl mx-auto">
      <form onSubmit={onSubmit} className="bg-zinc-800 p-10 rounded-lg mt-2">
        <label htmlFor="title" className="text-white">
          Title
        </label>
        <input
          type="text"
          id="title"
          name="title"
          {...register("title", { required: true })}
          className="bg-zinc-700 p-3 rounded-lg block w-full mb-3 text-white"
        />
        {errors.title && <span>This field is required</span>}

        <label htmlFor="description" className="text-white">
          Description
        </label>
        <textarea
          id="description"
          name="description"
          {...register("description", { required: true })}
          className="bg-zinc-700 p-3 rounded-lg block w-full text-white mb-3"
        />
        {errors.description && <span>This field is required</span>}

        {params.id ? (
          <button
            type="submit"
            className="bg-indigo-500 p-3 rounded-lg block w-full mt-3 text-white"
          >
            Update Task
          </button>
        ) : (
          <button
            type="submit"
            className="bg-indigo-500 p-3 rounded-lg block w-full mt-3 text-white"
          >
            Create Task
          </button>
        )}
      </form>

      {params.id && (
        <div className="flex justify-end">
          <button
            className="bg-red-500 p-3 rounded-lg w-48 mt-3 text-white"
            onClick={async () => {
              const accepted = window.confirm("Are you sure?");
              if (accepted) {
                await deleteTask(params.id);
                toast.success("Task Removed", {
                  position: "bottom-right",
                  style: {
                    background: "#101010",
                    color: "#fff",
                  },
                });
                navigate("/tasks");
              }
            }}
          >
            Delete
          </button>
        </div>
      )}
    </div>
  );
}

export default TaskFormPage;
