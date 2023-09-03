import { useState } from "react";
import axios from 'axios';

export default function AddUserForm() {
  const [name, setName] = useState("");

  const addUser = () => {
    // APIエンドポイントを適切に設定してください
    axios.post("http://localhost:8080/createusers", { name })
      .then(response => {
        if (response.status === 200) {
          alert("User added successfully!");
          // ここでユーザー一覧を更新する処理を書くこともできます
        }
      })
      .catch(error => {
        alert("Error adding user");
        console.error(error);
      });
  };

  return (
    <div>
      <h2>Add User</h2>
      <input 
        type="text"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <button onClick={addUser}>Add User</button>
    </div>
  );
}
