import axios from "axios";
import { useEffect, useState } from "react";

function App() {
  const [Todos, setTodos] = useState<[{ Id: string; Todo: string }]>([
    { Id: "a", Todo: "asd" },
  ]);
  const [input, setInput] = useState("");
  const [click, setClick] = useState(true);

  async function get() {
    let res = await (await axios.get("http://sample-backend-1:5000/get")).data;

    setTodos(res);
  }
  useEffect(() => {
    get();
  }, [click]);

  let handleChange = (e: React.FormEvent<HTMLInputElement>) => {
    setInput(e.currentTarget.value);
    console.log(input);
  };

  return (
    <>
      <div>
        <form>
          <input type="text" onChange={handleChange} />
          <button
            onClick={(e) => {
              e.preventDefault();
              setClick(!click);
              axios.post("http://backend:5000/add", {
                entry: input,
              });
            }}
          >
            Add Todo
          </button>
        </form>
        {Todos.map(({ Id, Todo }) => {
          console.log(Id);
          return <li>{Todo}</li>;
        })}
      </div>
    </>
  );
}

export default App;
