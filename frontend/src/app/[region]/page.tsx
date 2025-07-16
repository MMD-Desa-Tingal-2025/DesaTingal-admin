import { lora } from "@/libs/fonts";

export default async function RegionDetails({
  params,
}: {
  params: Promise<{ region: string }>;
}) {
  const region = (await params).region;
  return <div className={lora.className}>Region : {region}</div>;
}
