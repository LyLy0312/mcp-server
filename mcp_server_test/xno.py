import httpx
from mcp.server.fastmcp import FastMCP
import os
from dotenv import load_dotenv

load_dotenv()

mcp = FastMCP("xno_api_checker")

# Get credentials from .env
XNO_EMAIL = os.getenv("XNO_EMAIL")
XNO_PASSWORD = os.getenv("XNO_PASSWORD")

BASE_URL = "https://api-v2.xno.vn"

# List of private APIs that require login
XNO_PRIVATE_APIS = [
    "https://xno.vn/api/orders/open",
    "https://xno.vn/api/user/balance",
    "https://xno.vn/api/trades/recent", 
    f"{BASE_URL}/reports/v1/reports?page=0&limit=5",
    f"{BASE_URL}/ai-bot/v1/follows",
    f"{BASE_URL}/ta-bot/v1/VNINDEX/overview?page=0&limit=1000&resolution=D",
]

# Login and return authenticated client
async def login_xno() -> httpx.AsyncClient:
    client = httpx.AsyncClient(timeout=10)
    try:
        login_payload = {
            "email": XNO_EMAIL,
            "password": XNO_PASSWORD
        }
        resp = await client.post(f"{BASE_URL}/users/v1/auth/token", json=login_payload)
        if resp.status_code == 200 and "accessToken" in resp.json():
            token = resp.json()["accessToken"]
            client.headers.update({"Authorization": f"Bearer {token}"})
            print("[LOGIN SUCCESS] Token saved.")
        else:
            print("[LOGIN FAIL]", resp.text)
    except Exception as e:
        print("[ERROR LOGIN]", e)
    return client

# Check API status
async def fetch_xno_api_status(client: httpx.AsyncClient, url: str) -> dict:
    try:
        response = await client.get(url)
        return {
            "url": url,
            "status_code": response.status_code,
            "ok": response.status_code == 200
        }
    except httpx.HTTPError as e:
        return {
            "url": url,
            "status_code": None,
            "ok": False,
            "error": str(e)
        }

# Tool: Check all private APIs
@mcp.tool()
async def check_all_private_apis() -> str:
    """Login and check all private APIs."""
    client = await login_xno()
    results = []
    for url in XNO_PRIVATE_APIS:
        res = await fetch_xno_api_status(client, url)
        if res["ok"]:
            results.append(f"[OK] {url} → {res['status_code']}")
        else:
            results.append(f"[FAIL] {url} → {res.get('error', 'Status ' + str(res['status_code']))}")
    await client.aclose()
    return "\n".join(results)

# Tool: Count OK / FAIL
@mcp.tool()
async def count_private_api_status() -> str:
    """Login and count OK/FAIL APIs."""
    client = await login_xno()
    ok_count, fail_count = 0, 0
    for url in XNO_PRIVATE_APIS:
        res = await fetch_xno_api_status(client, url)
        if res["ok"]:
            ok_count += 1
        else:
            fail_count += 1
    await client.aclose()
    return f" OK: {ok_count} |  FAIL: {fail_count}"

if __name__ == "__main__":
    mcp.run(transport="stdio")
