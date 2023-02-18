import * as SingleGames from "./SingleGame.tsx";

interface GamesProps {
  games: SingleGames.Game[]
}

export default function Games({ games }: GamesProps) {
  return (
    <div class="flex flex-col gap-2 pt-2 w-full">
      {games?.map((game) => <SingleGames.default data={game}/>)}
    </div>
  );
}