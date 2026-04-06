import { Heading, Subtext } from "@/components/ui/Typography";

interface AuthHeaderProps {
  title: string;
  subtitle: string;
}

export default function AuthHeader({ title, subtitle }: AuthHeaderProps) {
  return (
    <div className="text-center space-y-2 mb-12">
      <Heading>{title}</Heading>
      <Subtext>{subtitle}</Subtext>
    </div>
  );
}
