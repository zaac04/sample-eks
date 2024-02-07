import axios from "axios";
import { useEffect, useState } from "react";

function App() {
  const [Todos, setTodos] = useState<[{ Id: string; Todo: string }]>([
    { Id: "a", Todo: "asd" },
  ]);
  const [input, setInput] = useState("");
  const [click, setClick] = useState(true);

  async function get() {
    
    let res = await (await axios.get("http://k8s-kubesyst-backend-219ac3e0df-917590178.eu-central-1.elb.amazonaws.com/get"));
    setTodos(res.data);
    console.log(res.data)
    
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
              axios.post("http://k8s-kubesyst-backend-219ac3e0df-917590178.eu-central-1.elb.amazonaws.com/add", {
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
