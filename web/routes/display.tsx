import { HandlerContext, Handlers, PageProps } from "$fresh/server.ts";
import Games from "../components/Games.tsx";
import {Game } from "../components/SingleGame.tsx"

interface ApiGamesProps {
  message: Array<Game>;
}

export const handler: Handlers = {
  async GET(_req, ctx) {
    const resp = await fetch("http://127.0.0.1:8080/games");
    if (!resp) {
      return new Response("Project not found", { status: 404 });
    }
    const aaa: ApiGamesProps = await resp.json();
    return ctx.render(aaa);
  },
};

export default function GamesDisplay({ data }: PageProps<ApiGamesProps>) {
  return (
    <div class="flex gap-2 w-full">
      <Games games={data.message}/>
    </div>
  );
}
