import ytdl from "@distube/ytdl-core";

async function downloadVideo(input: string, output: string) {
  const stream = ytdl(input, {
    filter: (format) => format.container === "mp4",
    quality: "highestvideo",
  });

  const writeStream = Deno.openSync(output, { write: true, create: true });

  for await (const chunk of stream) {
    writeStream.write(chunk);
  }

  writeStream.close();
  console.log("Download complete");

  return output;
}

if (import.meta.main) {
  const [url, output = "video.mp4"] = Deno.args;

  if (!url || !output) {
    console.error(
      "Usage: deno -RENW downloader.ts <url> [output]",
    );
    Deno.exit(1);
  }

  console.log(`Downloading video from ${url}...`);

  const videoPath = await downloadVideo(url, output);
  console.log(`Video downloaded to ${videoPath}`);
}
