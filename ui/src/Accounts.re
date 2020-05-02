open Data;

[@react.component]
let make = (~accounts, ~refreshAccounts, ~selectAccount) => {
  React.useEffect0(() => {
    refreshAccounts();
    None;
  });

  let refreshAccounts = _event => refreshAccounts();

  <div>
    <button onClick=refreshAccounts> {ReasonReact.string("Refresh")} </button>
    <ul>
      {accounts
       |> Array.map(account =>
            <li
              key={account.id} onClick={_event => selectAccount(account.id)}>
              {ReasonReact.string(account.name)}
            </li>
          )
       |> React.array}
    </ul>
  </div>;
};
