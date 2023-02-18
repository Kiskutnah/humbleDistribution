export interface SingleGameProps {
  data: Game
}

export interface Game {
  CreatedAt?: string;
  UpdatedAt?: string;
  DeletedAt?: string;
  ID?: string;
  Name?: string;
  Key?: string;
}

export default function SingleGame(singleGameData: SingleGameProps) {
  const data = singleGameData.data
  return (
    <table>
      <thead>
        <td>
          CreatedAt
        </td>
        <td>
          UpdatedAt
        </td>
        <td>
          DeletedAt
        </td>
        <td>
          ID
        </td>
        <td>
          Name
        </td>
      </thead>
      <tbody>
        <tr>
          <td>{data.CreatedAt} test</td>
          <td>{data.UpdatedAt} test</td>
          <td>{data.DeletedAt} test</td>
          <td>{data.ID} test</td>
          <td>{data.Name} test</td>
        </tr>
      </tbody>
    </table>
  );
}
