let filterTasksByStatus = (tasks, statusId) =>
  Array.to_list(tasks)
  |> List.filter((task: Data.task) => task.statusId == statusId)
  |> Array.of_list;

// TODO: Don't use these hardcoded values.
let statuses = [|
  "2470e4eb-9702-4769-b2e3-c3c410e6f258",
  "ac04d1b9-64a5-474c-9517-531cbc246b42",
  "b5082c3b-ba21-4bec-8e39-ef0d0c68cb6a",
|];

[@react.component]
let make = (~tasks) => {
  <div className="board">
    {statuses
     |> Array.map(status =>
          <Bucket key={status} tasks={filterTasksByStatus(tasks, status)} />
        )
     |> React.array}
  </div>;
};
