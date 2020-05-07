let baseUrl = "http://localhost:8080";

type account = {
  id: string,
  name: string,
  createdAt: string,
  updatedAt: string,
};

type accounts = array(account);

let emptyAccount = {id: "", name: "", createdAt: "", updatedAt: ""};

type association = {
  accountId: string,
  accountName: string,
  defaultProjectId: string,
};

type associations = array(association);

type user = {
  id: string,
  name: string,
  email: string,
  defaultAccountId: string,
  associations: array(association),
  createdAt: string,
  updatedAt: string,
};

type users = array(user);

let emptyUser = {
  id: "",
  name: "",
  email: "",
  defaultAccountId: "",
  associations: [||],
  createdAt: "",
  updatedAt: "",
};

type project = {
  id: string,
  name: string,
  description: string,
  accountId: string,
  createdAt: string,
  updatedAt: string,
};

type projects = array(project);

let emptyProject = {
  id: "",
  name: "",
  description: "",
  accountId: "",
  createdAt: "",
  updatedAt: "",
};

type listItem = {
  id: string,
  name: string,
};

type listItems = array(listItem);

let usersToListItems = (users: users): array(listItem) =>
  users |> Array.map((user: user) => {id: user.id, name: user.name});

let accountsToListItems = (accounts: accounts): array(listItem) =>
  accounts
  |> Array.map((account: account) => {id: account.id, name: account.name});

let associationsToListItems = (associations: associations): array(listItem) =>
  associations
  |> Array.map((association: association) =>
       {id: association.accountId, name: association.accountName}
     );

let withId = (id, fieldList) =>
  if (id != "") {
    [("id", Json.Encode.string(id))] @ fieldList;
  } else {
    fieldList;
  };

module Decode = {
  let account = (json): account =>
    Json.Decode.{
      id: json |> field("id", string),
      name: json |> field("name", string),
      createdAt: json |> field("createdAt", string),
      updatedAt: json |> field("updatedAt", string),
    };

  let accounts = (json): array(account) =>
    Json.Decode.(json |> array(account));

  let association = (json): association =>
    Json.Decode.{
      accountId: json |> field("accountId", string),
      accountName: json |> field("accountName", string),
      defaultProjectId: json |> field("defaultProjectId", string),
    };

  let associations = (json): array(association) =>
    Json.Decode.(json |> array(association));

  let user = (json): user =>
    Json.Decode.{
      id: json |> field("id", string),
      name: json |> field("name", string),
      email: json |> field("email", string),
      defaultAccountId: json |> field("defaultAccountId", string),
      associations: json |> field("associations", associations),
      createdAt: json |> field("createdAt", string),
      updatedAt: json |> field("updatedAt", string),
    };

  let users = (json): array(user) => Json.Decode.(json |> array(user));
  let project = (json): project =>
    Json.Decode.{
      id: json |> field("id", string),
      name: json |> field("name", string),
      accountId: json |> field("accountId", string),
      description: json |> field("description", string),
      createdAt: json |> field("createdAt", string),
      updatedAt: json |> field("updatedAt", string),
    };

  let projects = (json): array(project) =>
    Json.Decode.(json |> array(project));

  let id = (json): string => Json.Decode.string(json);
};

module Encode = {
  let account = (account: account) => {
    let fieldList = [("name", Json.Encode.string(account.name))];

    Json.Encode.(object_(withId(account.id, fieldList)));
  };

  let project = (project: project) => {
    let fieldList = [
      ("name", Json.Encode.string(project.name)),
      ("description", Json.Encode.string(project.description)),
      ("accountId", Json.Encode.string(project.accountId)),
    ];
    Json.Encode.(object_(withId(project.id, fieldList)));
  };
};

let listAccounts = callback =>
  Js.Promise.(
    Fetch.fetchWithInit(
      baseUrl ++ "/account",
      Fetch.RequestInit.make(~credentials=Fetch.Include, ()),
    )
    |> then_(Fetch.Response.json)
    |> then_(json => {
         json
         |> Decode.accounts
         |> (
           accounts => {
             callback(accounts);
             resolve();
           }
         )
       })
    |> ignore
  );

let saveAccount = (account: account, callback) =>
  Js.Promise.(
    Fetch.fetchWithInit(
      baseUrl ++ "/account",
      Fetch.RequestInit.make(
        ~method_=Fetch.Post,
        ~body=Fetch.BodyInit.make(Json.stringify(Encode.account(account))),
        ~credentials=Fetch.Include,
        (),
      ),
    )
    |> then_(Fetch.Response.json)
    |> then_(json =>
         json
         |> Decode.id
         |> (
           id => {
             callback(id);
             resolve();
           }
         )
       )
    |> ignore
  );

let listProjects = (accountId, callback) =>
  Js.Promise.(
    Fetch.fetchWithInit(
      baseUrl ++ "/project?accountId=" ++ accountId,
      Fetch.RequestInit.make(~credentials=Fetch.Include, ()),
    )
    |> then_(Fetch.Response.json)
    |> then_(json =>
         json
         |> Decode.projects
         |> (
           projects => {
             callback(projects);
             resolve();
           }
         )
       )
    |> ignore
  );

let saveProject = (project, callback) =>
  Js.Promise.(
    Fetch.fetchWithInit(
      baseUrl ++ "/project",
      Fetch.RequestInit.make(
        ~method_=Fetch.Post,
        ~body=Fetch.BodyInit.make(Json.stringify(Encode.project(project))),
        ~credentials=Fetch.Include,
        (),
      ),
    )
    |> then_(Fetch.Response.json)
    |> then_(json =>
         json
         |> Decode.id
         |> (
           id => {
             callback(id);
             resolve();
           }
         )
       )
    |> ignore
  );

let fetchToken = (userId, callback) =>
  Js.Promise.(
    Fetch.fetchWithInit(
      baseUrl ++ "/token/" ++ userId,
      Fetch.RequestInit.make(~credentials=Fetch.Include, ()),
    )
    |> then_(Fetch.Response.text)
    |> then_(_res => {
         callback();
         resolve();
       })
    |> ignore
  );

let fetchCurrentUser = callback =>
  Js.Promise.(
    Fetch.fetchWithInit(
      baseUrl ++ "/user",
      Fetch.RequestInit.make(~credentials=Fetch.Include, ()),
    )
    |> then_(Fetch.Response.json)
    |> then_(json => {
         json
         |> Decode.user
         |> (
           user => {
             callback(user);
             resolve();
           }
         )
       })
    |> ignore
  );
