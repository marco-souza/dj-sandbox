import { downloadVideo } from "./downloader.ts";

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
