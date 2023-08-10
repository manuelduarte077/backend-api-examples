import { Link } from "react-router-dom";

export default function Navigation() {
  return (
    <div>
      <Link to="/tasks">
        <h1>Tasks App</h1>
      </Link>
      <button>
        <Link to="/tasks-create">Create Task</Link>
      </button>
    </div>
  );
}
