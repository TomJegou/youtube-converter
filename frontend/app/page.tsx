import Image from "next/image";
import youtubeIcon from "@/public/youtube.png"

export default function Home() {
  return (
    <main className="min-h-screen flex flex-col justify-center items-center">
      <Image src={youtubeIcon} alt="youtube-icon" />
      <form action="/api" method="POST" className="flex flex-row flex-wrap gap-2">
        <input type="text" className="p-4 rounded-xl w-[50vw] outline-none" name="url" />
        <input type="submit" value="submit" className="p-4 bg-[#8EE3FB] text-black rounded-xl hover:scale-95" />
      </form>
    </main>
  );
}
