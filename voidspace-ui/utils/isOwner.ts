import type { User } from "../types";

type Entity =
  | { id?: string | number; username?: string }
  | string
  | number
  | null
  | undefined;

export function isOwner(entity: Entity, user: User | null): boolean {
  if (!entity || !user) return false;

  if (typeof entity === "string" || typeof entity === "number") {
    return entity.toString() === user.id.toString() || entity === user.username;
  }

  if (typeof entity === "object") {
    if (entity.id && entity.id.toString() === user.id.toString()) return true;
    if (entity.username && entity.username === user.username) return true;
  }

  return false;
}
