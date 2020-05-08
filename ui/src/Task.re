[@react.component]
let make = (~task: Data.task) => {
  <div className="task">
    <div className="title">
      {ReasonReact.string(task.title)}
    </div>
    <div>
      {ReasonReact.string(task.description)}
    </div>
  </div>;
};
