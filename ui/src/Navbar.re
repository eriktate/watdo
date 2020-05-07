[@react.component]
let make = (~currentUser: Data.user, ~selectedAccount, ~selectAccount) => {
  <nav>
    <div className="nav-left">
      <div className="banner"> {ReasonReact.string("WH")} </div>
      <Dropdown
        items={Data.associationsToListItems(currentUser.associations)}
        selectedId=selectedAccount
        select=selectAccount
      />
    </div>
    <ul>
      <li> <a href="#"> {ReasonReact.string("Dashboard")} </a> </li>
      <li> <a href="#tasks"> {ReasonReact.string("Tasks")} </a> </li>
      <li>
        <a href="#profile"> {ReasonReact.string(currentUser.name)} </a>
      </li>
    </ul>
    <div className="clearfix" />
  </nav>;
};
