export type TenantBootstrap = {
  tenant: {
    id: string;
    name: string;
    slug: string;
    domain: string;
    status: string;
    config: {
      theme?: {
        brandName?: string;
        primary?: string;
        secondary?: string;
        appearance?: string;
      };
    };
  };
};

export async function fetchTenantBootstrap(slug: string): Promise<TenantBootstrap> {
  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080"}/api/v1/tenant/bootstrap`,
    {
      headers: {
        "X-Tenant-Slug": slug
      },
      cache: "no-store"
    }
  );

  if (!response.ok) {
    const payload = (await response.json().catch(() => null)) as { error?: string } | null;
    throw new Error(payload?.error ?? "failed to bootstrap tenant");
  }

  return (await response.json()) as TenantBootstrap;
}
