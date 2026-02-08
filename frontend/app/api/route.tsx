import { NextRequest, NextResponse } from "next/server";

export async function POST(request: NextRequest, response: NextResponse) {
    let apiHostname = "localhost:46460"
    if (process.env.API_HOSTNAME) {
        apiHostname = process.env.API_HOSTNAME
    }
    const formData = await request.formData()
    const res = await fetch(`http://${apiHostname}/dl`, {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "url": formData.get("url")?.toString()
        })
    })
    if (!res.ok) {
        return NextResponse.json({"message": "error"})
    }
    return res
}