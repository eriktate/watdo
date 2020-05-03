type state = {
  accounts: Data.accounts,
  projects: Data.projects,
  loadingAccounts: bool,
  loadingProjects: bool,
  selectedAccount: option(string),
  selectedProject: option(string),
};

type action =
  | LoadingAccounts
  | LoadingProjects
  | LoadedAccounts(Data.accounts)
  | LoadedProjects(Data.projects)
  | SelectAccount(string)
  | SelectProject(string);

let initialState = {
  accounts: [||],
  projects: [||],
  loadingAccounts: true,
  loadingProjects: false,
  selectedAccount: None,
  selectedProject: None,
};

let refreshAccounts = (dispatch, ()) =>
  Data.listAccounts(payload => dispatch(LoadedAccounts(payload)));

let refreshProjects = (dispatch, accountId, ()) =>
  Data.listProjects(accountId, payload => dispatch(LoadedProjects(payload)));

let selectAccount = (dispatch, id) => {
  refreshProjects(dispatch, id, ());
  dispatch(SelectAccount(id));
};

let selectProject = (dispatch, id) => dispatch(SelectProject(id)) |> ignore;

let findAccountById = (accounts: Data.accounts, id) =>
  Array.to_list(accounts)
  |> List.find((account: Data.account) => account.id == id);

let findProjectById = (projects: Data.projects, id) =>
  Array.to_list(projects)
  |> List.find((project: Data.project) => project.id == id);

[@react.component]
let make = () => {
  let (state, dispatch) =
    React.useReducer(
      (state, action) =>
        switch (action) {
        | LoadingAccounts => {...state, loadingAccounts: true}
        | LoadingProjects => {...state, loadingProjects: true}
        | LoadedAccounts(accounts) => {
            ...state,
            loadingAccounts: false,
            selectedAccount: None,
            accounts,
          }
        | LoadedProjects(projects) => {
            ...state,
            loadingProjects: false,
            selectedProject: None,
            projects,
          }
        | SelectAccount(id) => {...state, selectedAccount: Some(id)}
        | SelectProject(id) => {...state, selectedProject: Some(id)}
        },
      initialState,
    );

  // initial load of accounts
  React.useEffect0(() => {
    refreshAccounts(dispatch, ());
    None;
  });

  let refreshAccounts = refreshAccounts(dispatch);
  let refreshProjects = refreshProjects(dispatch);
  let selectAccount = selectAccount(dispatch);
  let selectProject = selectProject(dispatch);
  let account =
    switch (state.selectedAccount) {
    | Some(id) => Some(findAccountById(state.accounts, id))
    | None => None
    };

  let project =
    switch (state.selectedProject) {
    | Some(id) => Some(findProjectById(state.projects, id))
    | None => None
    };

  <main>
    <Navbar />
    <AccountEdit account refreshAccounts />
    <Accounts accounts={state.accounts} refreshAccounts selectAccount />
    {switch (state.selectedAccount) {
     | Some(id) =>
       <div>
         <ProjectEdit project refreshProjects={() => refreshProjects(id)} />
         <Projects
           projects={state.projects}
           refreshProjects={refreshProjects(id)}
           selectProject
         />
       </div>
     | None => <div />
     }}
  </main>;
};
