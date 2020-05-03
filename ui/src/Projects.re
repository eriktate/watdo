[@react.component]
let make = (~projects: Data.projects, ~refreshProjects, ~selectProject) => {
  let refreshProjects = _event => refreshProjects();
  <div>
    <div> {ReasonReact.string("Projects")} </div>
    <button type_="button" onClick=refreshProjects>
      {ReasonReact.string("Refresh")}
    </button>
    <ul>
      {projects
       |> Array.map((project: Data.project) =>
            <li
              key={project.id} onClick={_event => selectProject(project.id)}>
              {ReasonReact.string(project.name)}
            </li>
          )
       |> React.array}
    </ul>
  </div>;
};
