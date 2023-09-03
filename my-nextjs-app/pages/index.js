import { useEffect, useState } from "react";
import axios from 'axios';

export default function Home() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    // APIエンドポイントを適切に設定してください
    axios.get("http://localhost:8888/users")
      .then(response => setUsers(response.data))
      .catch(error => console.error("Error fetching users:", error));
  }, []);

  return (
    <div>
      <h1>User List</h1>
      <ul>
        {users.map((user, index) => (
          <li key={index}>{user.name}</li>
        ))}
      </ul>
      <AddUserForm />
    </div>
  );
}
