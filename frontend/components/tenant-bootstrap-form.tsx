"use client";

import { useState, useTransition } from "react";

import { fetchTenantBootstrap, type TenantBootstrap } from "@/lib/api";

export function TenantBootstrapForm() {
  const [slug, setSlug] = useState("demo");
  const [result, setResult] = useState<TenantBootstrap | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isPending, startTransition] = useTransition();

  function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setError(null);

    startTransition(async () => {
      try {
        const payload = await fetchTenantBootstrap(slug);
        setResult(payload);
      } catch (submitError) {
        setResult(null);
        setError(submitError instanceof Error ? submitError.message : "request failed");
      }
    });
  }

  return (
    <div className="panel hero-side bootstrap-card">
      <h2>Tenant Bootstrap</h2>
      <p>
        Gunakan slug tenant untuk mengambil konfigurasi awal sebelum login. Default seed backend
        saat ini adalah <strong>demo</strong>.
      </p>

      <form onSubmit={handleSubmit}>
        <label htmlFor="tenant-slug">
          Slug tenant
          <input
            id="tenant-slug"
            name="tenant-slug"
            value={slug}
            onChange={(event) => setSlug(event.target.value)}
            placeholder="mis. smpn1punggur"
          />
        </label>

        <button type="submit" disabled={isPending}>
          {isPending ? "Memuat..." : "Ambil konfigurasi tenant"}
        </button>
      </form>

      {error ? <div className="tenant-result">Error: {error}</div> : null}

      {result ? (
        <div className="tenant-result">
          <strong>{result.tenant.name}</strong>
          <div>Slug: {result.tenant.slug}</div>
          <div>Domain: {result.tenant.domain}</div>
          <div>Status: {result.tenant.status}</div>
        </div>
      ) : null}
    </div>
  );
}
