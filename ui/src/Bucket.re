[@react.component]
let make = (~tasks) => {
  <div className="bucket">
    {tasks |> Array.map((task: Data.task) => <Task key={task.id} task />) |> React.array}
  </div>;
};
