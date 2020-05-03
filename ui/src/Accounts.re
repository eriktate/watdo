[@react.component]
let make = (~accounts: Data.accounts, ~refreshAccounts, ~selectAccount) => {
  React.useEffect0(() => {
    refreshAccounts();
    None;
  });

  let refreshAccounts = _event => refreshAccounts();

  <div>
    <div> {ReasonReact.string("Accounts")} </div>
    <button onClick=refreshAccounts> {ReasonReact.string("Refresh")} </button>
    <ul>
      {accounts
       |> Array.map((account: Data.account) =>
            <li
              key={account.id} onClick={_event => selectAccount(account.id)}>
              {ReasonReact.string(account.name)}
            </li>
          )
       |> React.array}
    </ul>
  </div>;
};
