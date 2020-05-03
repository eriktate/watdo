type state = {
  saving: bool,
  project: Data.project,
};

type action =
  | SetProject(Data.project)
  | SetName(string)
  | SetDescription(string)
  | Saving
  | Saved
  | Clear;

let initialState = project => {
  let initProject =
    switch (project) {
    | Some(proj) => proj
    | None => Data.emptyProject
    };
  {saving: false, project: initProject};
};

let saveProject = (dispatch, project, _event) => {
  dispatch(Saving);
  Data.saveProject(
    project,
    payload => {
      Js.log2("Saved project: ", payload);
      dispatch(Saved);
    },
  );
};

[@react.component]
let make = (~project, ~refreshProjects) => {
  let (state, dispatch) =
    React.useReducer(
      (state, action) =>
        switch (action) {
        | SetProject(project) => {...state, project}
        | SetName(name) => {
            ...state,
            project: {
              ...state.project,
              name,
            },
          }
        | SetDescription(description) => {
            ...state,
            project: {
              ...state.project,
              description,
            },
          }
        | Saving => {...state, saving: true}
        | Saved =>
          refreshProjects();
          {saving: false, project: Data.emptyProject};
        | Clear => {...state, project: Data.emptyProject}
        },
      initialState(project),
    );

  // update project data when new props are passed in
  React.useEffect1(
    () => {
      switch (project) {
      | Some(project) =>
        dispatch(SetProject(project));
        None;
      | None => None
      }
    },
    [|project|],
  );

  // event handlers
  let saveProject = saveProject(dispatch);
  let setName = event => {
    let name = ReactEvent.Form.target(event)##value;
    dispatch(SetName(name));
  };
  let setDescription = event => {
    let description = ReactEvent.Form.target(event)##value;
    dispatch(SetDescription(description));
  };
  let clear = _event => dispatch(Clear);

  <form>
    {if (state.project.id == "") {
       ReasonReact.string("Create new project");
     } else {
       ReasonReact.string(
         "Editing " ++ state.project.name ++ ": " ++ state.project.id,
       );
     }}
    <div className="fieldSet">
      <label>
        {ReasonReact.string("Project Name:")}
        <input
          value={state.project.name}
          type_="input"
          name="projectName"
          onChange=setName
        />
      </label>
    </div>
    <div className="fieldSet">
      <label>
        {ReasonReact.string("Project Description:")}
        <input
          value={state.project.description}
          type_="textarea"
          name="projectDescription"
          onChange=setDescription
        />
      </label>
    </div>
    <button type_="button" onClick={saveProject(state.project)}>
      {ReasonReact.string("Save")}
    </button>
    <button type_="button" onClick=clear>
      {ReasonReact.string("Clear")}
    </button>
  </form>;
};
