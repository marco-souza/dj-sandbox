import ytdl from "@distube/ytdl-core";

export async function downloadVideo(input: string, output: string) {
  const stream = ytdl(input, {
    filter: "audioonly",
    quality: "highestaudio",
  });

  const writeStream = Deno.openSync(output, { write: true, create: true });

  for await (const chunk of stream) {
    writeStream.write(chunk);
  }

  writeStream.close();
  console.log("Download complete");

  return output;
}
