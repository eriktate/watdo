// temporary value: user ID to emulate
let forcedUserId = "d163193e-6f6a-4a71-92a1-c76d3148559a";

type state = {
  currentUser: Data.user,
  accounts: Data.accounts,
  projects: Data.projects,
  tasks: Data.tasks,
  loadingUser: bool,
  loadingAccounts: bool,
  loadingProjects: bool,
  loadingTasks: bool,
  selectedAccount: string,
  selectedProject: option(string),
};

type action =
  | LoadingAccounts
  | LoadingProjects
  | LoadedUser(Data.user)
  | LoadedAccounts(Data.accounts)
  | LoadedProjects(Data.projects)
  | LoadedTasks(Data.tasks)
  | SelectAccount(string)
  | SelectProject(string);

let initialState = {
  currentUser: Data.emptyUser,
  accounts: [||],
  projects: [||],
  tasks: [||],
  loadingUser: true,
  loadingAccounts: true,
  loadingProjects: false,
  loadingTasks: false,
  selectedAccount: "",
  selectedProject: None,
};

let fetchUser = (dispatch, ()) =>
  Data.fetchCurrentUser(payload => dispatch(LoadedUser(payload)));

let refreshAccounts = (dispatch, ()) =>
  Data.listAccounts(payload => dispatch(LoadedAccounts(payload)));

let refreshProjects = (dispatch, accountId, ()) =>
  Data.listProjects(accountId, payload => dispatch(LoadedProjects(payload)));

let refreshTasks = (dispatch, projectId, ()) =>
  Data.listTasks(projectId, payload => dispatch(LoadedTasks(payload)));

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
        | LoadedUser(user) => {
            ...state,
            currentUser: user,
            selectedAccount: user.defaultAccountId,
          }
        | LoadingAccounts => {...state, loadingAccounts: true}
        | LoadingProjects => {...state, loadingProjects: true}
        | LoadedAccounts(accounts) => {
            ...state,
            loadingAccounts: false,
            selectedAccount: "",
            accounts,
          }
        | LoadedProjects(projects) => {
            ...state,
            loadingProjects: false,
            selectedProject: None,
            projects,
          }
        | LoadedTasks(tasks) => {...state, loadingTasks: false, tasks}
        | SelectAccount(id) => {...state, selectedAccount: id}
        | SelectProject(id) => {...state, selectedProject: Some(id)}
        },
      initialState,
    );

  // initial load of accounts
  React.useEffect0(() => {
    Data.fetchToken(forcedUserId, _payload =>
      Data.fetchCurrentUser(user => {
        dispatch(LoadedUser(user));
        refreshTasks(
          dispatch,
          (Array.to_list(user.associations)
          |> List.find((assoc: Data.association) =>
               assoc.accountId == user.defaultAccountId
             )).defaultProjectId,
          (),
        );
      })
    );
    None;
  });

  let refreshAccounts = refreshAccounts(dispatch);
  let refreshProjects = refreshProjects(dispatch);
  let selectAccount = selectAccount(dispatch);
  let selectProject = selectProject(dispatch);

  <>
    <Navbar
      currentUser={state.currentUser}
      selectedAccount={state.selectedAccount}
      selectAccount
    />
    <main>
      <Board tasks={state.tasks} />
    </main>
  </>;
};
